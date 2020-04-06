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
  getUser(id, type): any {
    let query = qs.stringify({id, type});
    return this.http.get<any>(`api/getUser?${query}`)
  }

}
