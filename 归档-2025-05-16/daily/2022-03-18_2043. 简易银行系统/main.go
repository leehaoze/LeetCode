package main

type Bank struct {
	accountCount int
	money        []int64
}

func Constructor(balance []int64) Bank {
	return Bank{
		money:        balance,
		accountCount: len(balance),
	}
}

func (this *Bank) accountIsValid(account int) bool {
	return account >= 1 && account <= this.accountCount
}

func (this *Bank) moneyEnough(account int, money int64) bool {
	return this.money[account-1] >= money
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if this.accountIsValid(account1) && this.accountIsValid(account2) && this.moneyEnough(account1, money) {
		this.money[account1-1] -= money
		this.money[account2-1] += money
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if this.accountIsValid(account) {
		this.money[account-1] += money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if this.accountIsValid(account) && this.moneyEnough(account, money) {
		this.money[account-1] -= money
		return true
	}
	return false
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
