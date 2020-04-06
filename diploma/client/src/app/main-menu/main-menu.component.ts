import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-main-menu',
  templateUrl: './main-menu.component.html',
  styleUrls: ['./main-menu.component.css']
})
export class MainMenuComponent implements OnInit {

  constructor(
    public activatedRoute: ActivatedRoute,
    private apiService: ApiService,

  ) { }

  public items

  ngOnInit(): void {
    this.activatedRoute.params.subscribe(async ({ type }) => {
      let user = sessionStorage.getItem('user')
      this.apiService.getMenuItems(user).subscribe(res => {
        let menu = res.menu
        menu.forEach(el => {
          if (el.sysname == type) {
            this.items = el.child
          }
        })
      });
    })
}
  goto(url) {

  }


}
