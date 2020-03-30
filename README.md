# Flutter + Centrifugo Notification Demo

Code are copied from : 

1. https://github.com/centrifugal/centrifuge-dart/blob/master/example/flutter_app/lib/main.dart
2. https://github.com/nitishk72/flutter_app_local_notification/blob/master/lib/main.dart

## Dependencies

1. Centrifugo [Install here](https://github.com/centrifugal/centrifugo#how-to-install)
2. Golang [Install here](https://golang.org/dl/)

## How to run ?

1. git clone https://github.com/codenoid/flutter-centrifugo-push-notification
2. run centrifugo with config `centrifugo --client_insecure --config=centrifugo-config.json`
3. Customize centrifugo url (flutter and fcpn_server)
4. cd to `fcpn_server` then `go run .`
5. run your flutter app, Connect then subscribe
6. send your notification via fcpn_server (point 4)
7. congrats, you have push notification without FCM
