def convert_to_decimal(num: int):
    """Converts a binary number to decimal using the doubling method.

    Args:
        num (int): The binary number to convert.

    Returns:
        int: The resulting decimal number.
    """

    result: int = 0
    base: int = 2
    current_number: int

    for n in str(num):
        current_number = int(n)
        total = result * base + current_number
        result = total

    return total


print(convert_to_decimal(10000))
