const fs = require('fs');

readThenSortToTarget('./input.txt', './output.txt')

async function readThenSortToTarget(source, target) {
    try {
        const data = await readFile(source)
        var arr = data.toString().split(',')
        arr.sort(function(o1, o2) {
            return o1 - o2;
        })
        await writeFile(target, arr.join(','))
    } catch(e) {
        console.log(e.message)
    }
}

function readFile(source) {
    return new Promise(function(resolve, reject) {
        fs.readFile(source, function(err, data) {
            if(err) {
                reject(err)
            } else {
                resolve(data)
            }
        })
    })
}

function writeFile(target, data) {
    return new Promise(function(resolve, reject){
        fs.writeFile(target, data, function(err, data){
            if(err) {
                reject(err)
            } else {
                resolve(data)
            }
        })
    })
}
