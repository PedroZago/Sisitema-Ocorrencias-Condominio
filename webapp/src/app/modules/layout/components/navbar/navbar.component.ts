import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { SairService } from '../../services';
import { UsuarioService } from '../../../../services';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  usuario_nome!: string;
  usuario_id!: number;
  url_foto_perfil!: any;
  authenticated: number = 0;
  regex = /([\S]+?(?!\S))/g;

  constructor(
    private usuariosService: UsuarioService,
    private saireService: SairService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.usuariosService.buscarUsuario()
      .subscribe(
        response => {
          this.usuario_nome = response.nome.match(this.regex)['0'];
          this.url_foto_perfil = `http://localhost:3000/api/usuario/foto-perfil/${response.id}`;
          this.authenticated = 1;
        },
        () => {
          this.authenticated = 2;
        }
      );


  }

  deslogarUsuario(): void {
    this.saireService.deslogarUsuario()
      .subscribe(
        () => {
          this.router.navigate(['/entrar']);
        }
      )
  }

}
