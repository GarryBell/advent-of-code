// This function was defined alongside the sheet
function IS_20(x, y) {
    let i = x
    while (i < y + 1) {
        if ((i - 20) % 40 === 0) {
            return true
        }
        i += 1
    }
    return false
}