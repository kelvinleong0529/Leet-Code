function reverseWords(s: string): string {
  let ans = "";
  let temp = "";

  for (let i = 0; i < s.length; i++) {
    if (s[i] === " ") {
      ans += temp.split("").reverse().join("");
      ans += " ";
      temp = "";
    } else {
      temp += s[i];
    }
  }

  return ans + temp.split("").reverse().join("");
};