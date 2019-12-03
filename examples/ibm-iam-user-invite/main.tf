provider "ibm" {}

resource "ibm_iam_user_invite" "invite_user" {
    users = ["${var.user1}", "${var.user2}"]
}

