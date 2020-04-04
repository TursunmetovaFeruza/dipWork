import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  constructor(
    private apiService: ApiService,
    private router: Router
  ) { }
public items
  ngOnInit(): void {
    let user = sessionStorage.getItem('user')
    this.apiService.getMenuItems(user).subscribe(res=>{
      let menu = res.menu
      menu.forEach(el => {
       if (el.userinfo == user) {
         this.items = el.child
       } 
      })
    })
  }

  goto(url) {

  }

}
