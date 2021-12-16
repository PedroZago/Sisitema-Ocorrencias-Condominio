import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { LoginService } from '../services';
import { LoginModel } from '../models';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  login!: LoginModel;

  loginForm = new FormGroup({
    email: new FormControl(null, [
      Validators.required,
      Validators.email
    ]),
    senha: new FormControl(null, [
      Validators.required,
      Validators.minLength(8)
    ])
  });

  constructor(
    private loginService: LoginService,
    private router: Router
  ) { }

  ngOnInit(): void {
  }

  get email(): any {
    return this.loginForm.get('email');
  }

  get senha(): any {
    return this.loginForm.get('senha');
  }

  doLogin(): void {
    this.login = this.loginForm.value;

    this.loginService.doLogin(this.login)
      .subscribe(
        message => {
          //console.log(message);
          this.router.navigate(['/ocorrencias']);
        },
        error => { }
      )
  }

}
