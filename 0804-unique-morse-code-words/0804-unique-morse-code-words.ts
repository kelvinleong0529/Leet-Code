const MORSE: string[] = [
    ".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---",
    "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-",
    "...-", ".--", "-..-", "-.--", "--.."
];

function uniqueMorseRepresentations(words: string[]): number {
    const hashMap: { [key: string]: boolean } = {};
    for (const word of words) {
        const morseRepresentation = getMorseRepresentation(word);
        hashMap[morseRepresentation] = true;
    }
    return Object.keys(hashMap).length;
}

function getMorseRepresentation(word: string): string {
    let answer = "";
    for (const char of word) {
        const asciiValue = charToAscii(char);
        answer += MORSE[asciiValue];
    }
    return answer;
}

function charToAscii(char: string): number {
    return char.charCodeAt(0) - 'a'.charCodeAt(0);
}