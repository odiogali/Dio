# Smart Review
Smart Review is a project that seeks to automate the process of deciding what to review. It deals with:
- Deciding what note to review based on what you've reviewed in the current cycle
- Formatting the notes such that they can be viewed as they are in your markdown editor
- Sending an email notification informing you that you need to review your notes as well as the IP and port which the server is on

# How to Use It
1. Run this command in your terminal `git clone https://github.com/odiogali/goSmartReviewer.git`
2. Create a `.env` file in the root project directory which will hold the following: the email address the message should be sent to (`TO=...`), the email address the message will be sent from (specified `FROM=...`), and the email password or app password for the email address specified in `FROM`.
3. Then, run the command `go run format.go main.go [list_of_directories]` where `[list_of_directories]` refers to the directories you want the program to choose the notes from all separated by spaces
6. Find a machine you can run the web server on / AWS or some other cloud computing service but ensure that this device has access to the your markdown notes
**NOTE**: If the terminal is displaying errors involving the specified dependencies, you may need to run `go mod download <DEPENDENCY_NAME>` for github dependencies and `go get <DEPENDENCY_NAME>` for gopkg dependencies.
