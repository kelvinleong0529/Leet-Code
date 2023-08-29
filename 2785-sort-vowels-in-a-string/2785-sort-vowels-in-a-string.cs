public class Solution {
    public string SortVowels(string s) {
        string vowels = "aeiouAEIOU";
        int[] vowelIndices = new int[s.Length];
        char[] vowelChars = new char[s.Length];
        int vowelCount = 0;

        // Identify vowel indices and collect vowel characters
        for (int i = 0; i < s.Length; i++)
        {
            if (vowels.Contains(s[i]))
            {
                vowelIndices[vowelCount] = i;
                vowelChars[vowelCount] = s[i];
                vowelCount++;
            }
        }

        // Sort the vowel characters in non-decreasing order
        Array.Sort(vowelChars, 0, vowelCount);
        
        // Create the resulting string t
        char[] t = s.ToCharArray();
        for (int i = 0; i < vowelCount; i++)
        {
            t[vowelIndices[i]] = vowelChars[i];
        }

        return new string(t);
    }
}