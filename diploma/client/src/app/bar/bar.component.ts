import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-bar',
  templateUrl: './bar.component.html',
  styleUrls: ['./bar.component.css']
})
export class BarComponent implements OnInit {
  public url
  public menuItems

  constructor(
    private apiService: ApiService,
    private router: Router
  ) { 
    router.events.subscribe((val) => {
      this.url = router.url
    });
  }
  

  ngOnInit(): void {
    let user = sessionStorage.getItem('user')
    this.apiService.getMenuItems(user).subscribe(res => {
      this.menuItems = res.menu
      this.menuItems.forEach(el => {
        el.clicked = true
      })
    })
  }

  goto(sysname) {
    this.menuItems.forEach(el => {
      if (el.sysname == sysname) {
        el.clicked = el.clicked ? false : true
        this.router.navigate([el.url])
      } else {
        el.clicked = true
      }
    });
  }
  gotosub(url){
    this.router.navigate([url])
  }
  Logout() {
    sessionStorage.removeItem('user')
    location.reload()
  }

  EditProfile() {
    this.router.navigate(['bar/profile'])
  }
}
