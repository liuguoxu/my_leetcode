# 快排


def partition(arr: list, left: int, right: int) -> int:
    i = left
    for j in range(left, right):
        if arr[j] < arr[right]:
            arr[i], arr[j] = arr[j], arr[i]
            i += 1

    arr[i], arr[right] = arr[right], arr[i]
    print(i)
    return i


def quick_sort(arr: list, left: int, right: int):
    if left >= right:
        return

    pivot = partition(arr, left, right)
    quick_sort(arr, left, pivot - 1)
    quick_sort(arr, pivot + 1, right)


def main():
    arr = [2, 1, 4, 6, 9]

    quick_sort(arr, 0, len(arr) - 1)
    print(arr)


if __name__ == '__main__':
    main()
