print('===============JAVASCRIPT===============');
print('Count of rows in User collection: ' + db.User.count());

db.User.insert({ id: '1', username: 'admin', password: "psw", isadmin: true });

print('===============AFTER JS INSERT==========');
print('Count of rows in User collection: ' + db.User.count());

alltest = db.User.find();
while (alltest.hasNext()) {
  printjson(alltest.next());
}