{% for ecs in compute %}
  {% if ecs.Configuration.cluster %}
   # Create ECS Cluster
    resource "aws_ecs_cluster" "{{ecs.Configuration.cluster}}" {
      name = "{{ecs.Configuration.cluster}}"
    }
  {% endif %}

  {% if ecs.Configuration.task_definition %}
  # Create ECS Task Definition
  resource "aws_ecs_task_definition" "{{ecs.Configuration.task_definition}}" {
    family                   = "{{ecs.Configuration.task_definition}}"
  }
  {% endif %}
  # Create ECS Service
  resource "aws_ecs_service" "{{ecs.Configuration.name}}" {
    name            = "{{ecs.Configuration.name}}"
    {% if ecs.Configuration.cluster %}
      cluster         = aws_ecs_cluster.{{ecs.Configuration.cluster}}.id
    {% endif %}
    {% if ecs.Configuration.task_definition %}
      task_definition = aws_ecs_task_definition.{{ecs.Configuration.task_definition}}.arn
    {% endif %}
    desired_count   = {{ecs.Configuration.desired_count|default_if_none:"1"|integer}}

    launch_type     = "{{ecs.Configuration.launch_type|default_if_none:"EC2"}}" 
    {% if ecs.Configuration.iam_role %}
      iam_role = "{{ecs.Configuration.iam_role}}"
    {%- endif -%}

    {% if ecs.Configuration.load_balancer %}
      load_balancer {
        target_group_arn = "{{ecs.Configuration.load_balancer.target_group_arn}}"
        container_name   = "{{ecs.Configuration.load_balancer.container_name}}"
        container_port   = {{ecs.Configuration.load_balancer.container_port}}
      }

    {%- endif -%}

    {% if ecs.Configuration.placement_constraints %}
      placement_constraints {
        type       = "{{ecs.Configuration.placement_constraints.type}}"
        expression = "{{ecs.Configuration.placement_constraints.expression}}"
      }
    {% endif %}

    {% if ecs.Configuration.placement_strategies %}
      placement_strategies {
        type   = "{{ecs.Configuration.placement_strategies.type}}"
        field  = "{{ecs.Configuration.placement_strategies.field}}"
      }
    {%- endif -%}

    {% if ecs.Configuration.tags %}
      tags = {
        {% for tag in ecs.Configuration.tags %}
          "{{ tag.key }}" = "{{ tag.value }}"
        {% endfor %}
      }
    {%- endif -%}
  }
{% endfor %}
