# Formatting and scanning strings

## Writing Strings

![Writing strings](./assets/writing_strings.png)

## Formatting strings

![Formatting Strings](./assets/formatting_strings.png)

### Understanding the formatting verb

![](./assets/gen_purpose_formatting_verb.png)

### Formatting Arrays, slices and maps

![](./assets/formatting_array_slice_maps.png)

### Using the integer formatting verbs

![](./assets/using_integer_formatting_verbs.png)

### Using the floating point formatting verbs

![](./assets/using_floating_point_formatting_verb.png)

#### Formatting verb modifier

![](./assets/formatting_verb_modifier.png)

### Using the string and character formatting verbs

![](./assets/using_string_character_formatting_verbs.png)

### Using Boolean Formatting verbs

![](./assets/using_boolean_formatting_verb.png)

### Using Pointer Formatting verbs

![](./assets/using_pointer_formatting_verb.png)

## Scanning Strings

* The fmt package provides functions for scanning strings, which is the process of parsing strings that contain values separated by spaces. 

![](./assets/scanning_string_methods.png)

### Scanning in a slice

![](./assets/scanning_into_a_slice.png)

### Dealing with newline characters

By default, scanning treats newlines in the same way as spaces, acting as separators between values.

The Scan function doesn’t stop looking for values until after it has received the number it expects and the first press of the Enter key is treated as a separator and not the termination of the input.

### Using a Different String Source



