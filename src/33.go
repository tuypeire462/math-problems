public class RandomGo {
    public static void main(String[] args) {
        // Example Go code to play a game of "Rock, Paper, Scissors"
        System.out.println("Player chooses rock");
        if (playerChoice == "rock") {
            System.out.println("Computer chooses paper");
        } else if (playerChoice == "paper") {
            System.out.println("Computer chooses scissors");
        } else {
            System.out.println("You win!");
        }
    }

    public static String playerChoice = "rock";
}
