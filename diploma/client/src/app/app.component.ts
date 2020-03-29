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
    // await this.loadNewsItems();
    if (this.router.url === '/') {
      this.router.navigate(['login'])
    }
  }

  async loadNewsItems() {
    this.apiService.getUsers().subscribe(res => {
      console.log(res)
    });
  }

  // async addPost() {
  //   this.user ={
  //     id: this.id,
  //     name: this.name,
  //     address: this.address,
  //     age: this.age,
  //     created_at: this.created_at,
  //     updated_at: this.updated_at
  //   }
  //   this.apiService.createUsers(this.user).subscribe(res => {
  //     console.log(res)
  //   });
  //   this.apiService.getUsers().subscribe(res => {
  //     console.log(res)
  //   });
  // }
}
