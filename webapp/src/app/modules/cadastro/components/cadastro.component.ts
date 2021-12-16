import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { CadastroService } from '../services';
import { CadastroModel } from '../models';
import { LoginService } from '../../login';
import { UsuarioService } from '../../../services';

@Component({
  selector: 'app-cadastro',
  templateUrl: './cadastro.component.html',
  styleUrls: ['./cadastro.component.css']
})
export class CadastroComponent implements OnInit {

  cadastro!: CadastroModel;

  constructor(
    private cadastroService: CadastroService,
    private loginService: LoginService,
    private usuarioService: UsuarioService,
    private router: Router
  ) { }

  ngOnInit(): void {
  }

  cadastroForm = new FormGroup({
    nome: new FormControl(null, [
      Validators.required
    ]),
    email: new FormControl(null, [
      Validators.required,
      Validators.email
    ]),
    senha: new FormControl(null, [
      Validators.required,
      Validators.minLength(8)
    ]),
    confirmar_senha: new FormControl(null, [
      Validators.required,
      Validators.minLength(8)
    ]),
    foto: new FormControl(null)
  });

  get email(): any {
    return this.cadastroForm.get('email');
  }

  get nome(): any {
    return this.cadastroForm.get('nome');
  }

  get senha(): any {
    return this.cadastroForm.get('senha');
  }

  get confirmar_senha(): any {
    return this.cadastroForm.get('confirmar_senha');
  }

  get foto(): any {
    return this.cadastroForm.get('foto');
  }

  cadastrarUsuario(): void {
    if (this.senha.value === this.confirmar_senha.value) {

      this.cadastro = { email: this.email.value, nome: this.nome.value, senha: this.senha.value };

      let formData: any = new FormData();
      formData.append("foto", this.foto.value);

      this.cadastroService.cadastrarUsuario(this.cadastro).subscribe(response => {
        this.loginService.doLogin(this.cadastro).subscribe(() => {
          this.usuarioService.uploadFotoPerfil(response.id, formData).subscribe(() => {
            this.router.navigate(['/ocorrencias'])
          }, error => { console.log(error); });
        }, error => { console.log(error); });
      }, error => { console.log(error); });

    } else {
      alert("Senhas não coincidem.")
    }
  }

  selecionarImagen(event: any) {
    const imagem = event.target.files[0];
    this.cadastroForm.patchValue({ foto: imagem });
    this.cadastroForm.get('foto')?.updateValueAndValidity();
  }

}
