import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Component({
  selector: 'app-comms',
  templateUrl: './comms.component.html',
  styleUrls: ['./comms.component.css']
})
export class CommsComponent implements OnInit {

  // myWebSocket: WebSocketSubject = webSocket('ws://localhost:8000');
  data :string = "";

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.data = "SET in ngOnInit"
    this.http.get("http://localhost:8080", { responseType: 'text' })
      .subscribe((resp:string) => {
        console.log(resp);
        this.data = resp.toString();
      });
  }

  onPush() {
    console.log("Pushed it");
  }
}

