import { Component, OnInit } from '@angular/core';
import { ApiService } from './api.service';
import {  Router } from '@angular/router';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  
currentUrl
  constructor(
    private apiService: ApiService,
    private router: Router

  ) { 

    router.events.subscribe((val)=>{
      this.currentUrl = val;
    })
  }
  
  async ngOnInit() {
    let user = sessionStorage.getItem('user')
    if (user) {
      this.router.navigate(['bar'])
    } else if (this.router.url === '/') {
      this.router.navigate(['login'])
    }
  }

  async loadNewsItems() {
    this.apiService.getUsers().subscribe(res => {
      console.log(res)
    });
  }
}
