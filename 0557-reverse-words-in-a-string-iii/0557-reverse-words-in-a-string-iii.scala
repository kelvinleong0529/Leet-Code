object Solution {
    def reverseWords(s: String): String = {
      var ans = ""
      var temp = ""

      for (i <- s) {
        if (i == ' ') {
          ans += temp.reverse
          ans += " "
          temp = ""
        } else {
          temp += i
        }
      }

      ans + temp.reverse
    }
}