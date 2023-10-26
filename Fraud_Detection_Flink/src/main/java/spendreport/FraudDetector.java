package spendreport;

import org.apache.flink.api.common.state.ValueState;
import org.apache.flink.api.common.state.ValueStateDescriptor;
import org.apache.flink.api.common.typeinfo.Types;
import org.apache.flink.api.common.state.MapState;
import org.apache.flink.api.common.state.MapStateDescriptor;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.streaming.api.functions.KeyedProcessFunction;
import org.apache.flink.util.Collector;
import entity.Alert;
import entity.Transaction;

public class FraudDetector extends KeyedProcessFunction<Long, Transaction, Alert> {

	private static final long serialVersionUID = 1L;

	private static final double SMALL_AMOUNT = 10.00;
	private static final double LARGE_AMOUNT = 500.00;
	private static final long ONE_MINUTE = 60 * 1000;

	private transient ValueState<Boolean> flagState;
	private transient ValueState<Long> timerState;
	private transient MapState<String, Integer> accountTransactionCounts;

	@Override
	public void open(Configuration parameters) throws Exception {
		ValueStateDescriptor<Boolean> flagDescriptor = new ValueStateDescriptor<>(
				"flag",
				Types.BOOLEAN);
		flagState = getRuntimeContext().getState(flagDescriptor);

		ValueStateDescriptor<Long> timerDescriptor = new ValueStateDescriptor<>(
				"timer-state",
				Types.LONG);
		timerState = getRuntimeContext().getState(timerDescriptor);

		MapStateDescriptor<String, Integer> descriptor = new MapStateDescriptor<>(
				"accountTransactionCounts",
				String.class,   // Key type (account number)
				Integer.class  // Value type (transaction count)
		);
		accountTransactionCounts = getRuntimeContext().getMapState(descriptor);
	}

	@Override
	public void processElement(
			Transaction transaction,
			Context context,
			Collector<Alert> collector) throws Exception {

		Boolean lastTransactionWasSmall = flagState.value();

		// Check if the flag is set
		if (lastTransactionWasSmall != null) {
			if (transaction.getAmount() > LARGE_AMOUNT) {
				//Output an alert downstream
				System.out.println("small");
				Alert alert = new Alert();
				alert.setId(transaction.getAccountId());

				collector.collect(alert);
			}
			// Clean up our state
			cleanUp(context);
		}

		if (transaction.getAmount() < SMALL_AMOUNT) {
			// set the flag to true
			//System.out.println("small");
			flagState.update(true);

			long timer = context.timerService().currentProcessingTime() + ONE_MINUTE;
			context.timerService().registerProcessingTimeTimer(timer);

			timerState.update(timer);
		}

		// Transaction count
		String accountNumber = String.valueOf(context.getCurrentKey());
		Integer currentCount = accountTransactionCounts.get(accountNumber);
		if (currentCount == null) {
			currentCount = 0;
		}

		// Check if the transaction count exceeds 5
		if (currentCount < 5) {
			currentCount++;
			accountTransactionCounts.put(accountNumber, currentCount);
			System.out.println("Account: " + accountNumber + ", Transaction Count: " + currentCount + " Amount " + transaction.getAmount());
		} else {
			// Output an alert downstream if the transaction count exceeds 5
			Alert alert = new Alert();
			alert.setId(transaction.getAccountId());
			collector.collect(alert);
		}
	}

	@Override
	public void onTimer(long timestamp, OnTimerContext ctx, Collector<Alert> out) throws Exception {
		// Remove flag and state after 1 minute
		timerState.clear();
		flagState.clear();
		accountTransactionCounts.clear();
	}

	private void cleanUp(Context ctx) throws Exception {
		// Delete timer
		Long timer = timerState.value();
		ctx.timerService().deleteProcessingTimeTimer(timer);

		// Clean up flag and state
		timerState.clear();
		flagState.clear();
		accountTransactionCounts.clear();
	}
}



