"""add model stage

Revision ID: 0a65ca497350
Revises: 3c495c9f691e
Create Date: 2022-06-12 11:52:44.995917

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = "0a65ca497350"
down_revision = "3c495c9f691e"
branch_labels = None
depends_on = None


def upgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    with op.batch_alter_table("model", schema=None) as batch_op:
        batch_op.add_column(sa.Column("recommended_stage", sa.Integer(), nullable=True))
    op.create_table(
        "model_stage",
        sa.Column("id", sa.Integer(), autoincrement=True, nullable=False),
        sa.Column("name", sa.String(length=100), nullable=True),
        sa.Column("map", sa.Float(), nullable=True),
        sa.Column("timestamp", sa.Float(), nullable=True),
        sa.Column("model_id", sa.Integer(), nullable=False),
        sa.Column("is_deleted", sa.Boolean(), nullable=False),
        sa.Column("create_datetime", sa.DateTime(), nullable=False),
        sa.Column("update_datetime", sa.DateTime(), nullable=False),
        sa.PrimaryKeyConstraint("id"),
    )
    op.create_index(op.f("ix_model_stage_name"), "model_stage", ["name"], unique=False)
    op.create_index(op.f("ix_model_stage_id"), "model_stage", ["id"], unique=False)
    op.create_index(
        op.f("ix_model_stage_model_id"), "model_stage", ["model_id"], unique=False
    )
    with op.batch_alter_table("model_stage", schema=None) as batch_op:
        batch_op.create_unique_constraint('uq_modelstage_name', ['model_id', 'name'])
    # ### end Alembic commands ###


def downgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    with op.batch_alter_table("model", schema=None) as batch_op:
        batch_op.drop_column("recommended_stage")
    with op.batch_alter_table("model_stage", schema=None) as batch_op:
        batch_op.drop_constraint("uq_modelstage_name", type_="unique")
    op.drop_index(op.f("ix_model_stage_model_id"), table_name="model_stage")
    op.drop_index(op.f("ix_model_stage_id"), table_name="model_stage")
    op.drop_index(op.f("ix_model_stage_name"), table_name="model_stage")
    op.drop_table("model_stage")
    # ### end Alembic commands ###
