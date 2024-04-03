package budget

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/nozzlium/heymat_backend/lib"
)

type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

type BudgetPlanResponse struct {
	ID              uint64       `json:"id"`
	Title           string       `json:"title"`
	User            UserResponse `json:"user"`
	Amount          uint64       `json:"amount"`
	AmountString    string       `json:"amountString"`
	Expense         uint64       `json:"expense"`
	ExpenseString   string       `json:"expenseString"`
	Balance         uint64       `json:"balance"`
	BalanceString   string       `json:"balanceString"`
	Private         bool         `json:"private"`
	CreatedAt       time.Time    `json:"createdAt"`
	CreatedAtString string       `json:"createdAtString"`
	UpdatedAt       time.Time    `json:"updatedAt"`
	UpdatedAtString string       `json:"updatedAtString"`
	HasBeenEdited   bool         `json:"hasBeenEdited"`
}

type BudgetPlanListResponse struct {
	PageNo      uint64               `json:"pageNo"`
	PageSize    uint64               `json:"pageSize"`
	RecordCount uint64               `json:"recordCount"`
	BudgetPlans []BudgetPlanResponse `json:"budgetPlans"`
}

func findAllBudgetPlanItems(
	ctx context.Context,
	DB *sql.DB,
	param BudgetPlanParam,
) (BudgetPlanListResponse, error) {
	queryRecordCount := `
    select
      count(id)
    from budget_plan
    where budget_plan.user_id = $1
      and budget_plan.title ilike '%' || $2 || '%'
  `
	var recordCount uint32
	err := DB.QueryRowContext(ctx, queryRecordCount, param.BudgetPlan.UserID, param.BudgetPlan.Title).
		Scan(&recordCount)
	if err != nil {
		log.Println(err)
		return BudgetPlanListResponse{}, err
	}

	if (recordCount - ((param.PageNo - 1) * param.PageSize)) < 0 {
		return BudgetPlanListResponse{
			PageNo: uint64(
				param.PageNo,
			),
			PageSize: 0,
			RecordCount: uint64(
				recordCount,
			),
			BudgetPlans: []BudgetPlanResponse{},
		}, nil
	}

	queryGetItems := `
    select
      budget_plan.id,
      budget_plan.title,
      user_account.id,
      user_account.username,
      budget_plan.amount,
      sum(expense.amount) as expense,
      budget_plan.private,
      budget_plan.created_at,
      budget_plan.updated_at
    from budget_plan
      join user_account on user_account.id = budget_plan.user_id
      left join expense on budget_plan.id = expense.budget_id
    where budget_plan.user_id = $1
      and budget_plan.deleted_at is null 
      and budget_plan.title ilike '%' || $2 || '%' 
    group by (user_account.id, budget_plan.id) 
    order by budget_plan.created_at desc
    offset $3
    limit $4;
  `
	rows, err := DB.QueryContext(
		ctx,
		queryGetItems,
		param.BudgetPlan.UserID,
		param.BudgetPlan.Title,
		(param.PageNo-1)*param.PageSize,
		param.PageSize,
	)
	if err != nil {
		log.Println(err)
		return BudgetPlanListResponse{}, err
	}
	defer rows.Close()

	budgetPlans := make(
		[]BudgetPlanResponse,
		0,
		param.PageSize,
	)
	for rows.Next() {
		budgetPlanResponse := *new(BudgetPlanResponse)
		mapRowToBudgetPlanResponse(
			&budgetPlanResponse,
			rows.Scan,
		)
		budgetPlans = append(
			budgetPlans,
			budgetPlanResponse,
		)
	}

	return BudgetPlanListResponse{
		PageNo: uint64(
			param.PageNo,
		),
		PageSize: uint64(
			len(budgetPlans),
		),
		RecordCount: uint64(
			recordCount,
		),
		BudgetPlans: budgetPlans,
	}, nil
}

func findBudgetPlanItemById(
	ctx context.Context,
	DB *sql.DB,
	id uint64,
) (BudgetPlanResponse, error) {
	query := `
    select
      budget_plan.id,
      budget_plan.title,
      user_account.id,
      user_account.username,
      budget_plan.amount,
      sum(expense.amount) as expense,
      budget_plan.private,
      budget_plan.created_at,
      budget_plan.updated_at
    from budget_plan
      join user_account on user_account.id = budget_plan.user_id
      left join expense on budget_plan.id = expense.budget_id
    where budget_plan.id = $1 
      and budget_plan.deleted_at is null
    group by (budget_plan.id, user_account.id)
  `

	budgetPlanResponse := *new(BudgetPlanResponse)
	row := DB.QueryRowContext(
		ctx,
		query,
		id,
	)
	err := mapRowToBudgetPlanResponse(
		&budgetPlanResponse,
		row.Scan,
	)
	if err != nil {
		log.Println(err)
		return BudgetPlanResponse{}, err
	}

	return budgetPlanResponse, nil
}

func mapRowToBudgetPlanResponse(
	budgetPlanResponse *BudgetPlanResponse,
	scan func(...any) error,
) error {
	var expense sql.NullInt64
	err := scan(
		&budgetPlanResponse.ID,
		&budgetPlanResponse.Title,
		&budgetPlanResponse.User.ID,
		&budgetPlanResponse.User.Username,
		&budgetPlanResponse.Amount,
		&expense,
		&budgetPlanResponse.Private,
		&budgetPlanResponse.CreatedAt,
		&budgetPlanResponse.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	if expense.Valid {
		budgetPlanResponse.Expense = uint64(
			expense.Int64,
		)
	}
	budgetPlanResponse.convertNumberDataToString()

	return nil
}

func (budgetPlanResponse *BudgetPlanResponse) convertNumberDataToString() {
	budgetPlanResponse.AmountString = lib.IntToCurrency(
		int64(
			budgetPlanResponse.Amount,
		),
	)
	budgetPlanResponse.ExpenseString = lib.IntToCurrency(
		int64(
			budgetPlanResponse.Expense,
		),
	)
	budgetPlanResponse.Balance = budgetPlanResponse.Amount - budgetPlanResponse.Expense
	budgetPlanResponse.BalanceString = lib.IntToCurrency(
		int64(
			budgetPlanResponse.Balance,
		),
	)

	budgetPlanResponse.HasBeenEdited = !budgetPlanResponse.CreatedAt.Equal(
		budgetPlanResponse.UpdatedAt,
	)
	budgetPlanResponse.CreatedAtString = lib.GetIdDateStringFull(
		budgetPlanResponse.CreatedAt,
	)
	budgetPlanResponse.UpdatedAtString = lib.GetIdDateStringFull(
		budgetPlanResponse.UpdatedAt,
	)
}
