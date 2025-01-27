Before running tests:

1.Please set up Rancher Server
2.Update default password.

Script update:
Navigate to rancherlogin.test_test.go file
update the user name and password.

Running tests perform below action:
1.Go mod tidy
2.Go mod vendor
3.ginkgo -v
