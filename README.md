
**prerequisites**
Go lanuage should be install on your machine.
url : https://go.dev/dl/
**Steps to Run  test :**
1.Clone this Repo into local machine
**Need to udpate user credential as per your environment:**
1. Navigate to below file in testscripts folder
   "rancherlogin.test_test.go"
2. Update the rancher url and credentials
3. Open termial and navigate automation folder
3.Execute this command "Go mod tidy"
4.Execute this command "Go mod vendor"
5.Navigate to testscripts folder from command line
6.execute this command "ginkgo -v" 
