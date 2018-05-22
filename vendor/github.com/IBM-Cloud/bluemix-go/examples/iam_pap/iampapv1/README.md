# IAMPAP example

This example shows how to assign an IAM policy to an user.

This creates a IAM policy and assigns it to user, list the policies of the user and deletes the policy created.

Example:

```go run main.go -org myOrg -space dev -user_id IBMid-2700015GTF -service_name my-service -role crn:v1:bluemix:public:iam::::role:Viewer```