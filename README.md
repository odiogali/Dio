# Smart Review (Written in Go)
Smart Review is a project that seeks to automate the process of deciding what to review. It deals with:
- Deciding what note to review based on what you've reviewed in the current cycle
- Formatting the notes such that they can be viewed as they are in your markdown editor
- Sending the note as a formatted email so you get reminded to review and know what to review

# How to Use It
1. Run this command in your terminal `git clone https://github.com/MisterBra1n/goSmartReview.git`
2. Start by specifying the three subjects you would like to review by replacing the paths to their directories
3. Specify the email you want the message to be sent from and the password denoted by 'mailPass'
4. Specify the email you want the message to be received at
5. You may also change the frequency you receive the review notes depending on how you use it
6. Find a machine you can run the code in perpetuity on / AWS or some other cloud computing service
**NOTE**: If the terminal is displaying errors involving the specified dependencies, you may need to run `go mod download <DEPENDENCY_NAME>` for github dependencies and `go get <DEPENDENCY_NAME>` for gopkg dependencies.

Once everything has been setup and the program has been run, you should receive a formatted HTML email that looks like this, once a day. ![image](https://github.com/MisterBra1n/goSmartReviewer/assets/108496802/a5b12fdf-38b6-4674-8b58-61a99180d57f)
