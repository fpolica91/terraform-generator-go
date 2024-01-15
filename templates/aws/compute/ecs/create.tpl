{% for ecs in compute %}
  resource "aws_ecs_service" "{{ecs.Configuration.name}}" {
    name = "{{ecs.Configuration.name}}"
    {% if ecs.Configuration.cluster %}
      cluster = "${aws_ecs_cluster.{{ecs.Configuration.cluster}}.id}"
    {%-endif-%}

    {% if ecs.Configuration.task_definition %}
      task_definition = "${aws_ecs_task_definition.{{ecs.Configuration.task_definition}}.arn}"
    {%-endif-%}

    {% if ecs.Configuration.desired_count %}
      desired_count = {{ecs.Configuration.desired_count | integer}}
    {%-endif-%}

    {% if ecs.Configuration.iam_role %}
      iam_role = "${aws_iam_role.{{ecs.Configuration.iam_role}}.arn}"
    {%-endif-%}

    {% if ecs.Configuration.load_balancer %}
      {% if ecs.Configuration.load_balancer.target_group_arn or  ecs.Configuration.container_name %}
        load_balancer {
          target_group_arn = "${aws_alb_target_group.{{ecs.Configuration.load_balancer.target_group_arn}}.arn}"
          container_name   = "{{ecs.Configuration.load_balancer.container_name}}"
          container_port   = {{ecs.Configuration.load_balancer.container_port}}
        }
      {% endif %}
    {%-endif-%}
    
    {% if ecs.Configuration.launch_type %}
      launch_type = "{{ecs.Configuration.launch_type}}"
    {%-endif-%}

    {% if ecs.Configuration.placement_constraints %}
      {% if ecs.Configuration.placement_constraints.type %}
          placement_constraints {
            {% if ecs.Configuration.placement_constraints.type %}
              type = "{{ecs.Configuration.placement_constraints.type}}"
            {%endif%}
            {% if ecs.Configuration.placement_constraints.expression %}
              expression = "{{ecs.Configuration.placement_constraints.expression}}"
            {%endif%}
          }
      {%- endif -%}
    {%-endif-%}

      {% if ecs.Configuration.placement_strategies %}
        {% if ecs.Configuration.placement_strategies.type %}
          placement_strategies {
            {% if ecs.Configuration.placement_strategies.type %}
              type = "{{ecs.Configuration.placement_strategies.type}}"
            {%endif%}
            {% if ecs.Configuration.placement_strategies.field %}
              field = "{{ecs.Configuration.placement_strategies.field}}"
            {%endif%}
          }
        {%- endif -%}
      {%-endif-%}


    {% if ecs.Configuration.tags %}
      tags = {
        {% for tag in ecs.Configuration.tags %}
          {{ tag.key }} = "{{ tag.value }}"
        {% endfor %}
      }
    {% endif %}
}

{% endfor %}