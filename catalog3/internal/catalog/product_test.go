package catalog

import (
	"context"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rickywinata/go-training/catalog3/internal/database/querier"
)

func TestCreateProduct(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input       *CreateProductInput
		want        *CreateProductOutput
		wantProduct *querier.Product
	}{
		"create product": {
			input: &CreateProductInput{Name: "product 1", Price: 5000},
			want: &CreateProductOutput{
				ProductView: &ProductView{
					Name:  "product 1",
					Price: 5000,
				},
			},
			wantProduct: &querier.Product{
				Name:  "product 1",
				Price: 5000,
			},
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := testDatabaseInstance.NewDatabase(t)
			q := querier.New(db)
			svc := NewService(db)

			gotOutput, err := svc.CreateProduct(ctx, tc.input)
			if err != nil {
				t.Fatal("failed to create product: ", err)
			}

			productInDB, err := q.FindProduct(ctx, gotOutput.Name)
			if err != nil {
				t.Fatalf("failed to find product: %s", err)
			}

			qt.Assert(t, gotOutput, qt.CmpEquals(), tc.want)
			qt.Assert(t,
				&productInDB,
				qt.CmpEquals(cmpopts.IgnoreFields(querier.Product{},
					"CreatedAt", "UpdatedAt")),
				tc.wantProduct)
		})
	}
}

func TestGetProduct(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		insertProductParams querier.InsertProductParams
		input               *GetProductInput
		want                *GetProductOutput
	}{
		"get a product": {
			insertProductParams: querier.InsertProductParams{
				Name:  "product 1",
				Price: 10000,
			},
			input: &GetProductInput{
				Name: "product 1",
			},
			want: &GetProductOutput{
				ProductView: &ProductView{
					Name:  "product 1",
					Price: 10000,
				},
			},
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := testDatabaseInstance.NewDatabase(t)
			q := querier.New(db)
			if _, err := q.InsertProduct(ctx, tc.insertProductParams); err != nil {
				t.Fatal(err)
			}
			svc := NewService(db)

			got, err := svc.GetProduct(ctx, tc.input)
			if err != nil {
				t.Fatal(err)
			}

			qt.Assert(t, got, qt.CmpEquals(), tc.want)
		})
	}
}
