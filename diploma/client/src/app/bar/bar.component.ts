import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-bar',
  templateUrl: './bar.component.html',
  styleUrls: ['./bar.component.css']
})
export class BarComponent implements OnInit {

  constructor(
    private apiService: ApiService,
    private router: Router
  ) { }
  public mainpage
  public menuItems
  ngOnInit(): void {
    let user = sessionStorage.getItem('user')
    this.apiService.getMenuItems(user).subscribe(res=>{
      this.menuItems = res.menu
      this.menuItems.forEach(el => {
        el.clicked = true
      })
      if (this.router.url == '/bar') {
        this.mainpage = true
      } else {
        this.mainpage = false
      }
    })
  }
goto(sysname){
  this.menuItems.forEach(el => {
    if (el.sysname == sysname) {
      el.clicked = el.clicked ? false : true
      this.router.navigate([el.url])
      console.log(this.router.url)
      this.mainpage = false
    }
  });

}
}
