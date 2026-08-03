## Manually Executing the Cloud Function

The SK Telegram bot is activated and executed automatically via Yandex.Cloud triggers; however, occasionally something goes wrong, and the bot must be executed manually so that the polls are sent to the respective chats.

To manually executed the poll, perform the following operations:

1. Log into Yandex.Cloud and navigate to Triggers.
2. Change both triggers' execution time to every 5 minutes.
3. Wait for up to 5 minutes until the trigger executes.
4. Once the polls are sent, revert the schedule to the original values (remove the backticks):
    - For the meal distribution poll: `0 16 ? * SUN`
      – For the driver's poll: `0 17 ? * SUN`
