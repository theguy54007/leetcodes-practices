=begin
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.



Example 1:

Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.
Example 2:

Input: n = 1, bad = 1
Output: 1
=end

# The is_bad_version API is already defined for you.
# @param {Integer} version
# @return {boolean} whether the version is bad
# def is_bad_version(version):
def first_bad_version(n)
    return n if n == 1
    start_at = 1
    end_at = n

    while start_at <= end_at
        mid = (start_at + end_at) / 2
        is_bad = is_bad_version(mid)
        before_is_bad = is_bad_version(mid - 1)

        if is_bad && !before_is_bad
            return mid
        elsif is_bad && before_is_bad
            end_at = mid - 1
        elsif !is_bad
            start_at = mid + 1
        end
    end

end
