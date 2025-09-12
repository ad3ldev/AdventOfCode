import java.io.File

fun main() {
    val filename = "input.txt"
    var count = 0
    
    try {
        val content = File(filename).readText()
        for (char in content) {
            when (char) {
                '(' -> count++
                ')' -> count--
            }
        }
        println(count)
    } catch (e: Exception) {
        println("Error: ${e.message}")
    }
}
