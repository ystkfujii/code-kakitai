package cart

import (
	"context"
	"time"

	"github.com/code-kakitai/go-pkg/errors"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	orderDomain "github/code-kakitai/code-kakitai/domain/order"
	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

type CartUseCase struct {
	cartRepo    cartDomain.CartRepository
	productRepo productDomain.ProductRepository
}

func NewCartUseCase(
	orderDomainService orderDomain.OrderDomainService,
	cartRepo cartDomain.CartRepository,
) *CartUseCase {
	return &CartUseCase{
		cartRepo: cartRepo,
	}
}

type CartUseCaseDto struct {
	ProductID string
	Count     int
}

func (uc *CartUseCase) Run(ctx context.Context, userID string, dtos []CartUseCaseDto, now time.Time) (string, error) {
	// カートから商品情報を取得
	cart, err := uc.cartRepo.FindByUserID(ctx, userID)
	if err != nil {
		return "", err
	}
	ids := make([]string, 0, len(cart.Products()))
	for _, cp := range cart.Products() {
		ids = append(ids, cp.ProductID())
	}

	// 在庫情報を取得
	products, err := uc.productRepo.FindByIDs(ctx, ids)
	if err != nil {
		return "", err
	}
	productMap := make(map[string]*productDomain.Product)
	for _, p := range products {
		productMap[p.ID()] = p
	}

	// カートに商品を追加
	for _, dto := range dtos {
		// dtoのcountが0の時はカートから商品を削除
		if dto.Count == 0 {
			if err := cart.RemoveProduct(dto.ProductID); err != nil {
				return "", err
			}
			continue
		}
		// 在庫の確認
		p, ok := productMap[dto.ProductID]
		if !ok {
			return "", errors.NewError("商品が見つかりません。")
		}
		if p.Consume(dto.Count); err != nil {
			return "", err
		}
		// 新しい商品の場合はカートに追加
		if err := cart.AddProduct(dto.ProductID, dto.Count); err != nil {
			return "", err
		}
	}
	// カートの永続化
	if err := uc.cartRepo.Save(ctx, cart); err != nil {
		return "", err
	}

	// 管理者とユーザーにメール送信
	return "nil", nil
}
