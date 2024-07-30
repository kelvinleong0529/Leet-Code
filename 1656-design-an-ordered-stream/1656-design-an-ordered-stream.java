class OrderedStream {

    String[] array;
    int pointer = 1;

    public OrderedStream(int n) {
        array = new String[n];
    }

    public List<String> insert(int idKey, String value) {
        array[idKey - 1] = value;
        List<String> answer = new ArrayList<>();
        while (pointer <= array.length && array[pointer - 1] != null) {
            answer.add(array[pointer - 1]);
            pointer++;
        }
        return answer;
    }
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * OrderedStream obj = new OrderedStream(n);
 * List<String> param_1 = obj.insert(idKey,value);
 */