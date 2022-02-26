import 'dart:html';

import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Jigree for Fun',
      home: Scaffold(
        appBar: AppBar(
          leading: IconButton(
            icon: const Icon(Icons.coffee_maker),
            onPressed: () {},
          ),
          title: const Text(
            "City Environment & Sustainability",
            textAlign: TextAlign.left,
            style: TextStyle(
              fontSize: 26,
              fontStyle: FontStyle.italic,
              color: Colors.white,
              fontWeight: FontWeight.bold,
            ),
          ),
          actions: [
            IconButton(
              icon: const Icon(Icons.fiber_dvr),
              onPressed: () {},
            ),
          ],
        ),
        body: ListView(
          children: [
            MyHomePage(
              title: "page",
            ),
          ],
        ),
        floatingActionButton: FloatingActionButton(
          child: const Icon(Icons.computer),
          foregroundColor: Colors.white70,
          backgroundColor: Colors.blueAccent,
          onPressed: () {},
        ),
      ),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  List<TableRow> rows = [];
  final myController = TextEditingController();

  _MyHomePageState({Key? key}) {
    rows.add(TableRow(children: <Widget>[
      Container(
        height: 32,
        color: Colors.white,
      ),
      Container(
        height: 32,
        color: Colors.white,
      ),
      Container(
        height: 32,
        color: Colors.white,
      ),
      Container(
        height: 32,
        color: Colors.white,
      ),
      Container(
        height: 32,
        color: Colors.white,
      ),
      Container(
        height: 32,
        color: Colors.white,
      ),
      Container(
        height: 32,
        color: Colors.white,
      ),
      Container(
        height: 32,
        color: Colors.white,
      ),
      Container(
        height: 32,
        color: Colors.white,
        child: TextFormField(
          controller: myController,
          style: const TextStyle(
            // fontSize: 24,
            color: Colors.blue,
            fontWeight: FontWeight.w600,
          ),
          decoration: const InputDecoration(
            border: UnderlineInputBorder(),
            hintText: 'Enter a name of city',
          ),
          showCursor: true,
          onEditingComplete: addRow,
        ),
      ),
      Container(
        height: 32,
        color: Colors.white,
        alignment: Alignment.centerLeft,
        child: TextButton(onPressed: addRow, child: const Icon(Icons.search)),
      ),
    ]));
  }

  void addRow() {
    setState(() {
      rows.add(TableRow(
        children: <Widget>[
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
                alignment: Alignment.center,
                height: 32,
                color: Colors.white,
                child: const Icon(Icons.train_sharp)),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.lightGreen,
              child: Text(myController.text),
            ),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.white,
              child: const Text("CO 2.5"),
            ),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.white,
              child: const Text("H 2.5"),
            ),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.white,
              child: const Text("NO2 2.5"),
            ),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.white,
              child: const Text("O3 2.5"),
            ),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.white,
              child: const Text("P 2.5"),
            ),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.white,
              child: const Text("PM10 2.5"),
            ),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.white,
              child: const Text("PM25 2.5"),
            ),
          ),
          TableCell(
            verticalAlignment: TableCellVerticalAlignment.middle,
            child: Container(
              alignment: Alignment.center,
              height: 32,
              color: Colors.white,
              child: const Text("SO2 2.5"),
            ),
          ),
        ],
      ));
    });
    myController.clear();

    print(rows.length);
  }

  // string co = 2;
  // string h = 3;
  // string no2 = 4;
  // string o3 = 5;
  // string p = 6;
  // string pm10 = 7;
  // string pm25 = 8;
  // string so2 = 9;
  @override
  Widget build(BuildContext context) {
    return Table(
      border: TableBorder.symmetric(),
      columnWidths: const <int, TableColumnWidth>{
        0: FixedColumnWidth(40),
        1: FlexColumnWidth(),
        2: FlexColumnWidth(),
        3: FlexColumnWidth(),
        4: FlexColumnWidth(),
        5: FlexColumnWidth(),
        6: FlexColumnWidth(),
        7: FlexColumnWidth(),
        8: FlexColumnWidth(),
        9: FlexColumnWidth(),
      },
      defaultVerticalAlignment: TableCellVerticalAlignment.middle,
      children: rows,
    );
  }
}
