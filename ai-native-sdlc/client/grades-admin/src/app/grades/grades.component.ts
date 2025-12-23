import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-grades',
  template: `
    <h2>Grades</h2>
    <ul>
      <li *ngFor="let g of grades">{{g.id}} - {{g.course}}: {{g.score}}</li>
    </ul>
  `
})
export class GradesComponent implements OnInit {
  grades: any[] = [];

  // inject HttpClient instead of using axios
  constructor(private http: HttpClient) {}

  ngOnInit() {
    // use Angular HttpClient (observables) instead of axios
    this.http.get<any[]>('http://localhost:8080/grades').subscribe(
      data => {
        this.grades = data;
      },
      err => {
        console.error('Failed to load grades', err);
      }
    );
  }
}
