# Dates, Times, and Durations

## Working with Dates and times

![](./assets/representing_date_and_time.png)

### Method for accessing time components

![](./assets/method_for_accessing_time_comp.png)

### Types used to describe time components

![](./assets/time_Components.png)

### Formatting Times as Strings

![](./assets/formatting_times_as_strings.png)

### The Layout Constants Defined by the time Package

![](./assets/reference_date_format.png)

### Parsing time values from strings

![](./assets/parsing_time_values_from_strings.png)

### Specifying a Parsing Location

* The Parse function assumes that dates and times expressed without a time zone are defined in Coordinated Universal Time (UTC). The ParseInLocation method can be used to specify a location that is used when no time zone is specified

### Function for creating Locations

![](./assets/functionForCreatingLocation.png)

### Embedding the time zone database

![](./assets/embedding_time_zone_db.png)

### Manipulating time values

![](./assets/method_for_working_with_time_values.png)

### Time duration representation

![](./assets/time_duration_representation.png)

#### Duration methods

![](./assets/duration_methods.png)

#### Creating Duration relative to time

![](./assets/creating_duration_relative_to_time.png)

#### Creating Duration from string

![](./assets/creating_duration_from_string.png)

## Using the Time Features for Goroutine and channels

![](./assets/function_of_time_package_for_goroutine_and_channels.png)

### Stopping and resetting timer

![](./assets/stopping_and_resetting_timer.png)

* Caution :- Be careful when stopping a timer. The timer’s channel is not closed, which means that reads from the channel will continue to block even after the timer has stopped

### Time Function for creating tick

![](./assets/time_function_for_creating_tick.png)