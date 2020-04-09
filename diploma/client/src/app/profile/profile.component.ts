import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  constructor(
    private apiService: ApiService,

  ) { }
  public user
  public isdisable = true
  ngOnInit(): void {
    let userType = sessionStorage.getItem('user')
    let id = sessionStorage.getItem('userId')
    this.apiService.getUser(id, userType).subscribe(res => {
      this.user = res.user
    });
    if (userType == 'admin') {
      this.isdisable = false
    } else if (userType == 'student') {
      // getUser(id, type) 
    } else if (userType == 'master') {

    }
  }

}
