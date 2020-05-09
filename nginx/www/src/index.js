import $ from "jquery";

var pulsEvents = []

var conn;
var hasWebsocket = false


if (window["WebSocket"]) {
    hasWebsocket = true
    conn = new WebSocket("ws://" + document.location.host + "/ws/")
}

if(hasWebsocket){
    setInterval(sendEventsToWs,5000)
}

function sendEventsToWs(){
    if(pulsEvents.length != 0){
        conn.send(JSON.stringify(pulsEvents))
        pulsEvents = []
    }
}

function addEvent(id, label){
    if(hasWebsocket){
        pulsEvents.push({id: id, label: label});
    }else{
        $.post('/api/',{ id: id, label: label})
            .done(function(data){
                if(data.status){
                    console.log(data.message)          
                }
            });
    }    
}

window.onbeforeunload = function (){
    if (hasWebsocket) {
        sendEventsToWs()
    }
}

if(hasWebsocket){
    conn.onmessage = function (evt) {
        console.log(evt.data)
    };
}

//bind events for send
$(function(){
    $('.events').on('click', function(){

        let id = $(this).attr('data-id');
        let label = $(this).attr('data-label')

        addEvent(id,label)

    })
})