package org.fi;

import java.io.Serializable;

public class Transaction implements Serializable {
    private long accountId;
    private double amount;

    public Transaction() {
    }

    public Transaction(long accountId, double amount) {
        this.accountId = accountId;
        this.amount = amount;
    }

    public long getAccountId() {
        return accountId;
    }

    public void setAccountId(long accountId) {
        this.accountId = accountId;
    }

    public double getAmount() {
        return amount;
    }

    public void setAmount(double amount) {
        this.amount = amount;
    }

    @Override
    public String toString() {
        return "Transaction{" +
                "accountId=" + accountId +
                ", amount=" + amount +
                '}';
    }
}
