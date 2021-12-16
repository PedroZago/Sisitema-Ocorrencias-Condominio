import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';

import { PerfilService } from '../services';
import { UsuarioService } from '../../../services';
import { PerfilModel } from '../models';

@Component({
  selector: 'app-perfil',
  templateUrl: './perfil.component.html',
  styleUrls: ['./perfil.component.css']
})
export class PerfilComponent implements OnInit {

  perfil!: PerfilModel;
  usuario_nome!: string;
  usuario_email!: string;
  usuario_id!: number;
  url_foto_perfil!: any;

  constructor(
    private usuarioService: UsuarioService,
    private perfilService: PerfilService
  ) { }

  ngOnInit(): void {
    this.usuarioService.buscarUsuario()
      .subscribe(
        response => {
          this.perfilForm.patchValue(response);
          this.url_foto_perfil = `http://localhost:3000/api/usuario/foto-perfil/${response.id}`;
          this.usuario_id = response.id;
          this.usuario_nome = response.nome;
          this.usuario_email = response.email;
        },
        error => { }
      );
  }

  perfilForm = new FormGroup({
    nome: new FormControl(null, [
      Validators.required
    ]),
    email: new FormControl(null, [
      Validators.required,
      Validators.email
    ]),
    foto: new FormControl(null)
  });

  get email(): any {
    return this.perfilForm.get('email');
  }

  get nome(): any {
    return this.perfilForm.get('nome');
  }

  get foto(): any {
    return this.perfilForm.get('foto');
  }

  atualizarUsuario(): void {
    this.perfil = { email: this.email.value, nome: this.nome.value, foto: this.foto.value };

    this.perfilService.atualizarUsuario(this.usuario_id, this.perfil)
      .subscribe(
        () => {
          window.location.reload();
        },
        error => { }
      );
  }

  selecionarImagen(event: any) {
    const imagem = event.target.files[0];
    this.perfilForm.patchValue({ foto: imagem });
    this.perfilForm.get('foto')?.updateValueAndValidity();

    let formData: any = new FormData();
    formData.append("foto", this.foto.value);

    this.usuarioService.uploadFotoPerfil(this.usuario_id, formData)
      .subscribe(
        () => { },
        error => { }
      );
  }

}
