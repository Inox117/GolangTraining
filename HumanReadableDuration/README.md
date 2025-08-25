# Human Readable Duration

## Task

Write a function which formats a duration, given as a number of seconds, in a human-friendly way.
The duration will be expressed as a combination of years, days, hours, minutes and seconds.

Additional Information:  
* The number of seconds must be a non-negative integer.
* If the number of seconds is zero, returns "now".  
  * A year is 365 days
  * A day is 24 hours
  * An hour is 60 minutes
  * A minute is 60 seconds
* Each time unit must be separated by a comma and a space (", ").  
* The last time unit must be separated by " and ".  
* Time units will be ordered from the bigger to the smallest.  
* No time units should be repeated
* A time unit should not appear if its value is zero

Example:  
* For seconds = 62, your function should return
  "1 minute and 2 seconds"
* For seconds = 3662, your function should return
  "1 hour, 1 minute and 2 seconds"

