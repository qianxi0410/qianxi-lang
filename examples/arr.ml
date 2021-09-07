let rev = fn(arr) {
    if len(arr) == 1 {
        return arr;
    } else {
        return push([last(arr)], rev(rest(arr)));
    }
};

rev([1, 2, 3]);