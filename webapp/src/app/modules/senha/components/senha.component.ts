import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';

import { SenhaService } from '../services';
import { UsuarioService } from '../../../services';
import { SenhaModel } from '../models';

@Component({
  selector: 'app-senha',
  templateUrl: './senha.component.html',
  styleUrls: ['./senha.component.css']
})
export class SenhaComponent implements OnInit {

  senha!: SenhaModel;
  usuario_id!: number;

  constructor(
    private usuariosService: UsuarioService,
    private senhaService: SenhaService
  ) { }

  ngOnInit(): void {
    this.usuariosService.buscarUsuario()
      .subscribe(
        response => {
          this.usuario_id = response.id;
        },
        error => { }
      );
  }

  SenhaForm = new FormGroup({
    senha_atual: new FormControl(null, [
      Validators.required,
      Validators.minLength(8)
    ]),
    senha_nova: new FormControl(null, [
      Validators.required,
      Validators.minLength(8)
    ]),
    confirmar_senha_nova: new FormControl(null, [
      Validators.required,
      Validators.minLength(8)
    ])
  });

  get senha_atual(): any {
    return this.SenhaForm.get('senha_atual');
  }

  get senha_nova(): any {
    return this.SenhaForm.get('senha_nova');
  }

  get confirmar_senha_nova(): any {
    return this.SenhaForm.get('confirmar_senha_nova');
  }

  atualizarSenha(): void {
    if (this.senha_nova.value === this.confirmar_senha_nova.value) {
      this.senha = { senha_atual: this.senha_atual.value, senha_nova: this.senha_nova.value };

      this.senhaService.atualizarSenha(this.usuario_id, this.senha)
        .subscribe(
          () => { window.location.reload(); },
          () => { }
        );
    } else {
      alert("Senhas novas não coincidem.")
    }
  }

}
