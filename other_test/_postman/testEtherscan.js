pm.test("Test log", function () {
    pm.response.to.have.status(200);
    // debugger
    var t = pm.response.text()
    var r = /<td>(?<idx>\d+)<\/td>.*?data-clipboard-text='(?<addr>\w{40,})'.*?<td>(?<num>[\d,\.]+)<\/td><td>(?<perc>[\d\.]+%)\s<div.*?(?<val>\$[\d\.,]{9,10})/img
    var tbl = [], list = [...t.matchAll(r)]
    for (var i in list){
        let gr = list[i].groups
        tbl[i] = `idx:${gr.idx} perc:${gr.perc} addr:${gr.addr} val:${gr.val} num:${gr.num}`
    }
    var first = list[0].groups.num, last = list[list.length-1].groups.num
    var str = `Num From ${last} To ${first}\n` + tbl.join("\n")
    console.clear()
    console.log(str)
});


/*
https://etherscan.io/token/generic-tokenholders2?a=0x9813037ee2218799597d83D4a5B6F3b6778218d9&sid=&m=light&s=230003022824847134336699287&p=1
*/