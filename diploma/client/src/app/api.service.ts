import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
declare var require: any
let qs = require('qs');
import * as _ from "lodash";



@Injectable()
export class ApiService {

  constructor(private http: HttpClient, ) {
  }
  getUsers() {
    return this.http.get<any>('api/users');
  }
  createUsers(value): any {
    return this.http.post<any>('api/users', value);
  }
  signIn(user): any {
    return this.http.post<any>('api/user', user)
  }
  getMenuItems(user): any {
    let query = qs.stringify({user});
    return this.http.get<any>(`api/getMenuItems?${query}`)
  }
  // getMessage201ByHour(startDate, finishDate) {
  //   let query = qs.stringify({startDate, finishDate});
  //   return this.http.get<any>(`/api/main/getMessage201ByHour?${query}`)
  // }

}
