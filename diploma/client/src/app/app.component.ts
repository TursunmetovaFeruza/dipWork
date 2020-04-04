import { Component, OnInit } from '@angular/core';
import { ApiService } from './api.service';
import {  Router } from '@angular/router';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  

  constructor(
    private apiService: ApiService,
    private router: Router

  ) { 

  }
  
  async ngOnInit() {
    let user = sessionStorage.getItem('user')
    localStorage.clear();
    if (user) {

      if (this.router.url === '/') {
        // this.router.navigate(['bar'])
      }
    } else if (this.router.url === '/') {
      this.router.navigate(['login'])
    }
  }

}
