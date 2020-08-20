package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

// Parameter store keys
var (
	KeyIxoFactor                      = []byte("ixoFactor")
	KeyClaimFeeAmount                 = []byte("ClaimFeeAmount")
	KeyEvaluationFeeAmount            = []byte("EvaluationFeeAmount")
	KeyNodeFeePercentage              = []byte("NodeFeePercentage")
	KeyEvaluationPayFeePercentage     = []byte("EvaluationPayFeePercentage")
	KeyEvaluationPayNodeFeePercentage = []byte("EvaluationPayNodeFeePercentage")
)

// payments parameters
type Params struct {
	IxoFactor                      sdk.Dec   `json:"ixo_factor" yaml:"ixo_factor"`
	ClaimFeeAmount                 sdk.Coins `json:"claim_fee_amount" yaml:"claim_fee_amount"`
	EvaluationFeeAmount            sdk.Coins `json:"evaluation_fee_amount" yaml:"evaluation_fee_amount"`
	NodeFeePercentage              sdk.Dec   `json:"node_fee_percentage" yaml:"node_fee_percentage"`
	EvaluationPayFeePercentage     sdk.Dec   `json:"evaluation_pay_fee_percentage" yaml:"evaluation_pay_fee_percentage"`
	EvaluationPayNodeFeePercentage sdk.Dec   `json:"evaluation_pay_node_fee_percentage" yaml:"evaluation_pay_node_fee_percentage"`
}

// ParamTable for payments module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(ixoFactor sdk.Dec, claimFeeAmount, evaluationFeeAmount sdk.Coins,
	nodeFeePercentage, evaluationPayFeePercentage,
	evaluationPayNodeFeePercentage sdk.Dec) Params {

	return Params{
		IxoFactor:                      ixoFactor,
		ClaimFeeAmount:                 claimFeeAmount,
		EvaluationFeeAmount:            evaluationFeeAmount,
		NodeFeePercentage:              nodeFeePercentage,
		EvaluationPayFeePercentage:     evaluationPayFeePercentage,
		EvaluationPayNodeFeePercentage: evaluationPayNodeFeePercentage,
	}

}

// default payments module parameters
func DefaultParams() Params {
	return Params{
		IxoFactor:                      sdk.OneDec(), // 1.0
		ClaimFeeAmount:                 sdk.NewCoins(sdk.NewInt64Coin(ixo.IxoNativeToken, 60000000)),
		EvaluationFeeAmount:            sdk.NewCoins(sdk.NewInt64Coin(ixo.IxoNativeToken, 40000000)),
		NodeFeePercentage:              sdk.NewDecWithPrec(5, 1), // 0.5
		EvaluationPayFeePercentage:     sdk.NewDecWithPrec(1, 1), // 0.1
		EvaluationPayNodeFeePercentage: sdk.NewDecWithPrec(2, 1), // 0.2
	}
}

// validate params
func ValidateParams(params Params) error {
	if params.IxoFactor.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter IxoFactor should be positive, is %s ", params.IxoFactor.String())
	}
	if params.ClaimFeeAmount.IsAnyNegative() {
		return fmt.Errorf("payments parameter ClaimFeeAmount should be positive, is %s ", params.ClaimFeeAmount.String())
	}
	if params.EvaluationFeeAmount.IsAnyNegative() {
		return fmt.Errorf("payments parameter EvaluationFeeAmount should be positive, is %s ", params.EvaluationFeeAmount.String())
	}
	if params.NodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter NodeFeePercentage should be positive, is %s ", params.NodeFeePercentage.String())
	}
	if params.EvaluationPayFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter EvaluationPayFeePercentage should be positive, is %s ", params.EvaluationPayFeePercentage.String())
	}
	if params.EvaluationPayNodeFeePercentage.LT(sdk.ZeroDec()) {
		return fmt.Errorf("payments parameter EvaluationPayNodeFeePercentage should be positive, is %s ", params.EvaluationPayNodeFeePercentage.String())
	}
	// TODO: validate according to param upper limits
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Payments Params:
  Ixo Factor:                               %s
  Claim Fee Amount:                         %s
  Evaluation Fee Amount:                    %s
  Node Fee Percentage:                      %s
  Evaluation Pay Fee Percentage:            %s
  Evaluation Pay Node Fee Percentage:       %s

`,
		p.IxoFactor, p.ClaimFeeAmount, p.EvaluationFeeAmount,
		p.NodeFeePercentage, p.EvaluationPayFeePercentage,
		p.EvaluationPayNodeFeePercentage,
	)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoFactor, &p.IxoFactor},
		{KeyClaimFeeAmount, &p.ClaimFeeAmount},
		{KeyEvaluationFeeAmount, &p.EvaluationFeeAmount},
		{KeyNodeFeePercentage, &p.NodeFeePercentage},
		{KeyEvaluationPayFeePercentage, &p.EvaluationPayFeePercentage},
		{KeyEvaluationPayNodeFeePercentage, &p.EvaluationPayNodeFeePercentage},
	}
}
