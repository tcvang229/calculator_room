

"use strict";

(function () {

    // setting up websocket connection
    //ws://aqueous-basin-43746.herokuapp.com/
    var websocket = new WebSocket("wss://" + window.location.host + "/ws");


    // setting up global variables needed

    var history_table = document.getElementById("history_table");
    var submit_value = document.getElementById("submit_value");

    var submit_button = document.getElementById("submit_button");
    submit_button.addEventListener("click", function () {
        var total_string = evaluateExpression(submit_value.value);
        if (total_string != null) {

            try {
                websocket.send(
                    JSON.stringify({
                        username: "client",
                        message: total_string,
                    })
                );
                console.log("sent!");
            } catch (error) {
                console.log(error.message);
            }
        }
    });

    var recent_submit = document.getElementById("recent_submit");



    // websocket handlers

    websocket.onerror = function (event) {
        console.log("error connecting: ", event);
    };

    websocket.onopen = function () {
        console.log("successful connection to server");

        updateHistoryTable(history_table);
    };

    websocket.onmessage = function (event) {
        var json_data = JSON.parse(event.data);
        recent_submit.innerHTML = json_data.message;
        console.log("received message: ", event);

        updateHistoryTable(history_table);
    };


    
    // functions that'll handle the HTML DOM for updating and deleting elements

    function removeTableRows(table) {
        try {
            while (true) {
                table.deleteRow(1)
            }
        } catch (error) {
            // will hit when no more row[1]
            console.log(error.message);
        }
    }

    function evaluateExpression(expression) {
        try {
            var equation = math.evaluate(expression);
            var total_string = submit_value.value + " = " + equation.toString();
            return total_string;
        } catch (error) {
            console.log(error.message);
            alert("Not a valid value, use '*' for multiplication, ex.: 2 * 19");
            return null;
        }
    }

    function updateHistoryTable(table) {
        try {
            var json_data = null;

            $.get("/getJson", function (data) {
                json_data = JSON.parse(data);
            }).fail(function () {
                console.log("failed with GET");
            }).done(function () {
                try {
                    removeTableRows(table);
                    var i;
                    if (json_data.length > 10) { i = json_data.length - 10; } else { i = 0; }

                    while (i < json_data.length) {
                        if (json_data[i].message != null || json_data[i].message != "") {
                            var row = table.insertRow(1);
                            var cell = row.insertCell(0);
                            cell.innerHTML = json_data[i].message;
                        }
                        i++;
                    }
                } catch (error) {
                    console.log(error.message);
                }
            });
        } catch (error) {
            console.log(error.message);
        }
    }
})();

