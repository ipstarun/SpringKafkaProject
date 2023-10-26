package com.example.demo.dto;

public class TransactionDTO {
	private long accountId;
	private double amount;


	//private long timestamp;

	@Override
	public String toString() {
		return "TransactionDTO [accountId=" + accountId + ". " + ", amount=" + amount + "]";
	}

	public long getAccountId() {
		return accountId;
	}

	public void setAccountId(long accountId) {
		this.accountId = accountId;
	}

//	public long getTimestamp() {
//		return timestamp;
//	}
//
//	public void setTimestamp(long timestamp) {
//		this.timestamp = timestamp;
//	}

	public double getAmount() {
		return amount;
	}

	public void setAmount(double amount) {
		this.amount = amount;
	}

	
}
