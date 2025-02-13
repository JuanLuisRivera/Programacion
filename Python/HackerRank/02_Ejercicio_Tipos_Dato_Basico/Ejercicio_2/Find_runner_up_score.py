if __name__ == '__main__':
    n = int(input())
    arr = map(int, input().split())
    
    arr_list = list(arr)
    max_list = max(arr_list)
    sorted_list = sorted(arr_list)
    
    for i in range(len(sorted_list)):
        if sorted_list[i] == max_list:
            print(sorted_list[i-1])
            break