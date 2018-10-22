const mongo = require('mongodb').MongoClient
const client = require('socket.io').listen(4000).sockets

//connect to mongo
mongo.connect('mongodb://127.0.0.1/mongoChat',(err,db)=>{
    if(err) {
        throw err
    }
    console.log('mongodb connected')
})