import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';

import { UnidadeModel } from '../models';
import { UnidadesService } from '../services';
import { UsuarioService } from '../../../services';
import { ResponsavelService, ResponsavelModel } from '../../responsavel';
import { BlocosService } from '../../blocos';
import { ApartamentosService } from '../../apartamentos';

@Component({
  selector: 'app-unidades',
  templateUrl: './unidades.component.html',
  styleUrls: ['./unidades.component.css']
})
export class UnidadesComponent implements OnInit {

  unidade!: UnidadeModel;
  responsavel_unidade!: ResponsavelModel;
  uni_selec: UnidadeModel = {};
  unidade_id!: number;
  usuario_id!: number;
  unidades: any = [];
  apartamentos: any = [];
  blocos: any = [];
  responsaveis: any = [];

  constructor(
    private unidadesService: UnidadesService,
    private apartamentosService: ApartamentosService,
    private responsavelService: ResponsavelService,
    private blocosService: BlocosService,
    private usuariosService: UsuarioService,
  ) { }

  ngOnInit(): void {
    this.buscarTodosApartamentos();
    this.buscarTodasUnidades();
    this.buscarTodosBlocos();
    this.buscarTodosUsuarios();
  }

  unidadeForm = new FormGroup({
    apartamento: new FormControl(null, [
      Validators.required
    ]),
    bloco: new FormControl(null, [
      Validators.required
    ]),
    responsavel: new FormControl(null, [
      Validators.required
    ])
  });

  unidadeModalForm = new FormGroup({
    apartamento_modal: new FormControl(null, [
      Validators.required
    ]),
    bloco_modal: new FormControl(null, [
      Validators.required
    ]),
    responsavel_modal: new FormControl(null, [
      Validators.required
    ])
  });

  get apartamento(): any {
    return this.unidadeForm.get('apartamento');
  }

  get bloco(): any {
    return this.unidadeForm.get('bloco');
  }

  get responsavel(): any {
    return this.unidadeForm.get('responsavel');
  }

  get apartamento_modal(): any {
    return this.unidadeModalForm.get('apartamento_modal');
  }

  get bloco_modal(): any {
    return this.unidadeModalForm.get('bloco_modal');
  }

  get responsavel_modal(): any {
    return this.unidadeModalForm.get('responsavel_modal');
  }

  cadastrarUnidade(): void {
    this.unidade = { apartamento_id: parseInt(this.apartamento.value), bloco_id: parseInt(this.bloco.value) };

    this.unidadesService.cadastrarUnidade(this.unidade)
      .subscribe(
        response => {
          this.responsavel_unidade = { unidade_id: parseInt(response.id), usuario_id: parseInt(this.responsavel.value) }
          this.responsavelService.cadastrarResponsavel(this.responsavel_unidade)
            .subscribe(
              response => {
                this.buscarTodasUnidades();
                this.unidadeForm.reset();
              },
              error => { }
            );
        },
        error => { }
      );
  }

  buscarTodasUnidades(): void {
    this.unidadesService.buscarTodasUnidades()
      .subscribe(
        response => {
          if (response != null) {
            this.unidades = [];
            response.forEach((unidade: any) => this.unidades.push(unidade));
          }
        },
        error => { }
      );
  }

  buscarUnidadeID(unidade: UnidadeModel): void {
    this.unidadesService.buscarUnidadeID(unidade.id!)
      .subscribe(
        response => {
          this.unidade_id = response.id;
          this.usuario_id = response.usuario_id;
          this.unidadeModalForm.patchValue({ apartamento_modal: response.apartamento_id, bloco_modal: response.bloco_id, responsavel_modal: response.usuario_id });
        },
        error => { }
      );
  }

  buscarTodosApartamentos(): void {
    this.apartamentosService.buscarTodosApartamentos()
      .subscribe(
        response => {
          this.apartamentos = [];
          response.forEach((apartamento: any) => this.apartamentos.push(apartamento));
        },
        error => { }
      );
  }

  buscarTodosUsuarios(): void {
    this.usuariosService.buscarTodosUsuarios()
      .subscribe(
        response => {
          this.responsaveis = [];
          response.forEach((usuario: any) => this.responsaveis.push(usuario));
        },
        error => { }
      );
  }

  buscarTodosBlocos(): void {
    this.blocosService.buscarTodosBlocos()
      .subscribe(
        response => {
          this.blocos = [];
          response.forEach((bloco: any) => this.blocos.push(bloco));
        },
        error => { }
      );
  }

  deletarUnidade(unidade: UnidadeModel): void {
    if (confirm(`Deseja deletar o unidade ${unidade.id}?`)) {
      this.unidadesService.deletarUnidade(unidade.id!)
        .subscribe(
          response => {
            this.responsavelService.deletarResponsavel(this.unidade_id, this.usuario_id)
              .subscribe(
                response => {
                  this.buscarTodasUnidades();
                },
                error => { }
              );
          },
          error => { }
        );
    }
  }

  atualizarUnidade(): void {
    this.uni_selec = { apartamento_id: parseInt(this.apartamento_modal.value), bloco_id: parseInt(this.bloco_modal.value) }

    this.unidadesService.atualizarUnidade(this.unidade_id, this.uni_selec)
      .subscribe(
        response => {
          this.responsavel_unidade = { unidade_id: this.unidade_id, usuario_id: parseInt(this.responsavel_modal.value) }
          this.responsavelService.atualizarResponsavel(this.unidade_id, this.usuario_id, this.responsavel_unidade)
            .subscribe(
              response => {
                this.buscarTodasUnidades();
                this.unidadeForm.reset();
              },
              error => { }
            );
        },
        error => { }
      );
  }

}
